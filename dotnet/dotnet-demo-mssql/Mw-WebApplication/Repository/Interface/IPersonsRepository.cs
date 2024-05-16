using WebApplication1.Model;

namespace WebApplication1.Repository.Interface
{
    public interface IPersonsRepository
    {
        Task<IEnumerable<Person>> GetAllPersonsData_SP();
        Task<IEnumerable<Person>> GetAllPersonsData_SQL();
    }
}
