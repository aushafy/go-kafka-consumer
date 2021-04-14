# execute kafka producer console, so we can send message to broker
  cd ..\..\..\..
  cd D:\Tools\kafka_2.13-2.7.0
  .\bin\windows\kafka-console-producer.bat --bootstrap-server localhost:9092 --topic message

# to start kafka broker
  cd ..\..\..\..
  cd D:\Tools\kafka_2.13-2.7.0
  .\bin\windows\kafka-server-start.bat .\config\server.properties

# RUN FIRST! before Kafka Broker, we need to start Zookeeper
  cd ..\..\..\..
  cd D:\Tools\kafka_2.13-2.7.0
  .\bin\windows\zookeeper-server-start.bat .\config\zookeeper.properties
