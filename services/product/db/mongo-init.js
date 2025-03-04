db.createUser({
  user: "superfast_delivery",
  pwd: "hxqpfreyim8bbde12z0cewjq",
  roles: [
    {
      role: "readWrite",
      db: "product_service_db",
    },
  ],
});
