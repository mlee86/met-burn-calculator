exports.up = function (knex, Promise) {
  return knex.schema.createTableIfNotExists("mets", function (table) {
    table.increments().primary();
    table.integer("met_code");
    table.float("met_units");
    table.string("Category");
    table.string("Description");
  });
};

exports.down = function (knex, Promise) {
  return knex.schema.dropTableIfExists("mets");
};
