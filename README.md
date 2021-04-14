# execute kafka producer console, so we can send message to broker
  1. `cd ..\..\..\..`
  2. `cd D:\Tools\kafka_2.13-2.7.0`
  3. `.\bin\windows\kafka-console-producer.bat --bootstrap-server localhost:9092 --topic message`

# to start kafka broker
  1. `cd ..\..\..\..`
  2. `cd D:\Tools\kafka_2.13-2.7.0`
  3. `.\bin\windows\kafka-server-start.bat .\config\server.properties`

# RUN FIRST! before Kafka Broker, we need to start Zookeeper
  1. `cd ..\..\..\..`
  2. `cd D:\Tools\kafka_2.13-2.7.0`
  3. `.\bin\windows\zookeeper-server-start.bat .\config\zookeeper.properties`
