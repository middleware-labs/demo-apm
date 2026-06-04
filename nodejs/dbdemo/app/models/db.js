const { Pool } = require("pg");
const dbConfig = require("../config/db.config.js");

const pool = new Pool({
  host: '1',
  port: 5432,
  user: 'postgres',
  password: 'postgres',
  database: 'todo'
});

// Without an 'error' listener, an unexpected drop on an idle pooled connection
// becomes an unhandled 'error' event that crashes the whole process.
pool.on("error", (err) => console.error("[db] idle client error:", err.message));

module.exports = pool;

// const mysql = require("mysql");
// const dbConfig = require("../config/db.config.js");

// const connection = mysql.createPool({
//   host: dbConfig.HOST,
//   user: dbConfig.USER,
//   password: dbConfig.PASSWORD,
//   database: dbConfig.DB
// });

// module.exports = connection;

