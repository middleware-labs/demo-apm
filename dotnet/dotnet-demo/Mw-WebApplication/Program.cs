using Middleware.APM;

var builder = WebApplication.CreateBuilder(args);
              
var configuration = new ConfigurationBuilder()
    .SetBasePath(Directory.GetCurrentDirectory())
    .AddJsonFile("appsettings.json", optional: false, reloadOnChange: true)
    .AddEnvironmentVariables()
    .Build();

builder.Services.ConfigureMWInstrumentation(configuration);

// builder.Logging.AddConfiguration(configuration.GetSection("Logging"));
// builder.Logging.AddConsole();
// builder.Logging.SetMinimumLevel(LogLevel.Debug);

// Add the following code after var app = builder.Build();
// Logger.Init(app.Services.GetRequiredService<ILoggerFactory>());

// Add services to the container.

builder.Services.AddControllers();
// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseHttpsRedirection();

app.UseAuthorization();

app.MapControllers();

app.Run();
