using WebApplication1.Model;
using WebApplication1.Data;
using WebApplication1.Repository.Interface;
using System.Data;
using Dapper;

namespace MWebApplication1.Repository
{
    public class PersonsRepository: IPersonsRepository
    {
        private readonly DapperContext _context;
        public PersonsRepository(DapperContext context)
        {
            this._context = context;
        }

        public async Task<IEnumerable<Person>> GetAllPersonsData()
        {
            using (var connection = this._context.CreateDBConnection())
            { 
                //Execute stored procedure and map the returned result to a Customer object  
                var response = await connection.QueryAsync<Person>("GetAllPersons", commandType: CommandType.StoredProcedure);
                return response;
            }
        }
    }
}
