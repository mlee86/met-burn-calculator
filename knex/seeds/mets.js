const seedData = require("../../seedData");
exports.seed = function (knex) {
  // Deletes ALL existing entries
  return (
    knex("mets")
      .del()
      // .then(() => knex("mets").del())
      .then(() => Promise.all([knex("mets").insert(seedData.data.mets)]))
      .then(() => {
        console.log("Data Seeded Successful");
      })
  );
};
