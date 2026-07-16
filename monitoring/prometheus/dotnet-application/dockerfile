FROM mcr.microsoft.com/dotnet/sdk:8.0 AS build

RUN mkdir -p /src/
COPY ./src/dotnet.csproj /src/dotnet.csproj

WORKDIR /src/
RUN dotnet restore

COPY ./src/ /src/
RUN mkdir /out/

RUN dotnet build dotnet.csproj --configuration Debug --no-restore
RUN dotnet publish dotnet.csproj --output /out --configuration Debug --no-restore

FROM mcr.microsoft.com/dotnet/aspnet:8.0 AS runtime

WORKDIR /app
COPY --from=build /out/ /app

ENTRYPOINT ["dotnet", "dotnet.dll"]