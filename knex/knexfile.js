// Update with your config settings.

module.exports = {
  development: {
    client: "postgresql",
    connection: {
      host: "127.0.0.1",
      database: "met_burn",
      password: "matt",
    },
  },

  staging: {
    client: "postgresql",
    connection: {
      database: "met_burn",
      user: "username",
      password: "password",
    },
    pool: {
      min: 2,
      max: 10,
    },
    migrations: {
      tableName: "knex_migrations",
    },
  },

  production: {
    client: "postgresql",
    connection: {
      database: "met_burn",
      user: "username",
      password: "password",
    },
    pool: {
      min: 2,
      max: 10,
    },
    migrations: {
      tableName: "knex_migrations",
    },
  },
};
