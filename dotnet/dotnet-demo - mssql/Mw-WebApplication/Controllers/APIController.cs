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

        [HttpGet("GetAllPersonsData")]
        public async Task<IActionResult> GetAllPersonsData()
        {
            var response = await _personRepository.GetAllPersonsData();
            if (!response.Any())
            {
                return NotFound("No Data Found");
            } 
            return Ok(response);
        }
    }
}
