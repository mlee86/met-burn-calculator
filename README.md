## MET Tracker

After cloning the repo, run `npm i` to install dependancies.

There are two database setups in this project.

### Json-Server

To start the json-server, cd into the server folder and run `npm i` again to install the json-server package.
Run the command `npm run server` or `yarn server` to start the database running on localhost:4000

### Postgres Database

Install postgres database on your machine
You will need to create a database called met_burn and you may have to set up a password on your user.
If you have to set up a password, change the password in knexfile.js in the development section to match what you set for your user
Run `cd knex`
Run `knex migrate:latest` to run the migration file and setup the table 'mets' in the database 'met_burn'
Run `knex seed: run` to seed the data into the table
