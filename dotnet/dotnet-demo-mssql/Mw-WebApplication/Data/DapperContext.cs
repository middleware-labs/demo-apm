using Microsoft.Data.SqlClient;
using System.Data;

namespace WebApplication1.Data
{
    public class DapperContext
    {
        private readonly IConfiguration _configuration;
        private readonly string connectionString;

        public DapperContext(IConfiguration configuration)
        {
            this._configuration = configuration;
            this.connectionString = _configuration.GetConnectionString("LocalDBConnectionString");
        }

        public IDbConnection CreateDBConnection() => new SqlConnection(connectionString);

    }
}
