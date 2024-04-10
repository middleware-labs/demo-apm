FROM mcr.microsoft.com/dotnet/aspnet:7.0 AS base

WORKDIR /app
RUN apt-get update && apt-get install -y curl unzip
RUN curl -sSfL https://install.middleware.io/apm/dotnet/v1.0.0-rc.1/scripts/mw-dotnet-auto-install.sh -O 
RUN bash mw-dotnet-auto-install.sh 
RUN rm -rf mw-dotnet-auto-install.sh
RUN chmod +x $HOME/.mw-dotnet-auto/instrument.sh

FROM mcr.microsoft.com/dotnet/sdk:7.0 AS build
WORKDIR /src
COPY Mw-WebApplication.csproj .
RUN dotnet restore Mw-WebApplication.csproj
COPY . .
WORKDIR /src
RUN dotnet build -c Release -o /app/build

FROM build AS publish
RUN dotnet publish -c Release -o /app/publish 

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .

CMD . $HOME/.mw-dotnet-auto/instrument.sh && dotnet Mw-WebApplication.dll
