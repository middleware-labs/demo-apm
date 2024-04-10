using Microsoft.AspNetCore.Mvc;
using WebApplication1.Repository.Interface;

namespace WebApplication1.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class APIController : ControllerBase
    {
        private readonly IPersonsRepository _personRepository;
        public APIController(IPersonsRepository personsRepository) 
        {
            this._personRepository = personsRepository;
        }

        [HttpGet("GetAllPersonsData_SP")]
        public async Task<IActionResult> GetAllPersonsData_SP()
        {
            var response = await _personRepository.GetAllPersonsData_SP();
            if (!response.Any())
            {
                return NotFound("No Data Found");
            } 
            return Ok(response);
        }
        
        [HttpGet("GetAllPersonsData_SQL")]
        public async Task<IActionResult> GetAllPersonsData_SQL()
        {
            var response = await _personRepository.GetAllPersonsData_SQL();
            if (!response.Any())
            {
                return NotFound("No Data Found");
            } 
            return Ok(response);
        }
    }
}
