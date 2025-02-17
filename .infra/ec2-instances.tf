resource "aws_instance" "bastion" {
  ami                         = var.ubuntu_ami
  instance_type               = "t3.micro"
  subnet_id                   = module.vpc.public_subnets[0]
  security_groups             = [aws_security_group.bastion_sg.id]
  key_name                    = aws_key_pair.bastion.key_name
  associate_public_ip_address = true

  user_data = <<-EOF
    #!/bin/bash
    set -e

    # Update and install Ansible
    sudo apt update -y
    sudo apt install -y ansible

    # Store the private key for SSH access to K8s nodes
    echo '${tls_private_key.k8s_key.private_key_pem}' > /home/ubuntu/k8s-key.pem
    chmod 600 /home/ubuntu/k8s-key.pem
    chown ubuntu:ubuntu /home/ubuntu/k8s-key.pem

    # Disable SSH strict host key checking for Ansible
    echo "StrictHostKeyChecking no" >> /etc/ssh/ssh_config
    echo "UserKnownHostsFile=/dev/null" >> /etc/ssh/ssh_config

    # Create Ansible inventory file
    touch /home/ubuntu/inventory.ini

    # Change ownership of inventory file
    chown ubuntu:ubuntu /home/ubuntu/inventory.ini
    chmod 644 /home/ubuntu/inventory.ini

    cat <<EOT > /home/ubuntu/inventory.ini
    [master]
    ${aws_instance.k8s_master.private_ip} ansible_user=ubuntu ansible_private_key_file=/home/ubuntu/k8s-key.pem ansible_ssh_common_args='-o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null'

    [workers]
    %{for instance in aws_instance.k8s_worker~}
    ${instance.private_ip} ansible_user=ubuntu ansible_private_key_file=/home/ubuntu/k8s-key.pem ansible_ssh_common_args='-o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null'
    %{endfor~}
  EOF

  tags = {
    Name = "bastion-host"
  }
}

resource "aws_instance" "k8s_master" {
  ami             = var.ubuntu_ami
  instance_type   = "t3.medium"
  subnet_id       = module.vpc.private_subnets[0]
  security_groups = [aws_security_group.k8s_control_plane_sg.id]
  key_name        = aws_key_pair.k8s.key_name

  tags = {
    Name = "k8s-master-01"
  }
}

resource "aws_instance" "k8s_worker" {
  count           = 2
  ami             = var.ubuntu_ami
  instance_type   = "t3.medium"
  subnet_id       = module.vpc.private_subnets[count.index]
  security_groups = [aws_security_group.k8s_data_plane_sg.id]
  key_name        = aws_key_pair.k8s.key_name

  tags = {
    Name = "k8s-worker-0${count.index + 1}"
  }
}
