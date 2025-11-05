Blightsanest - Stable Insights CLI

User Stories:
1) Allow users to apply known finance algorithms on the data (reasonable investments)
2) Allow users to compare the assets
3) Allow users to filter their assets (descending/increasing prices etc.)
4) Allow users to records their finances (net worth, owned assets, etc.) and to run simulations on them

3rd Party:
- Uses multiple PUBLIC APIs to get realtime data (listed below)

How It Works:
Runs third part APIs on the server to get data
Publishes the fetched data to clients
Runs database on another server
Publishes the database data to the clients
Clients run the desired operations on the data
Clients save the results to the database

### V0: - base application -
- Public crypto APIs
- Basic algorithms to follow the data
- Ability to compare multiple assets
- Ability to filter the assets on desired credentials

### Roadmap - V0
1) Server - Fetch and publish raw data with filtering ability
. Fetch live crypto data - to be published crypto feeds - from the public APIs
. Filter the data into the request struct
. Publish the filtered data

2) Client
. Subscribe to the publishing servers

### Structure - V0
. Server 
. Client

### V1: - first release -
- Authentication
- Database to record the assets, transactions

### Structure - V1
. Database
. Server
. Client

###V2:
- Simulations