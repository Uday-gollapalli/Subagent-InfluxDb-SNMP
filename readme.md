**Steps**

1. Run the subagent in the first assignment.
2. Create an Influxdb database named "anm".
3. Start the prober by following command which automatically logs the snmp responses into an Influxdb database "anm".
    
    go run prober.go <ip>:<community>:<port> <probing interval> <oid1> <oid2> ......<oidn>

    go run prober.go localhost:public:161 3 1.3.6.1.4.1.4171.40.2 1.3.6.1.4.1.4171.40.3

4. Import the dashboard.json file to grafana.

5. In the example json file, i used only two oids and they are logged into a measurement called "localhost" same as the ip in the above command.

