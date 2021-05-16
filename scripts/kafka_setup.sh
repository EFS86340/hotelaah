#! /bin/bash
#
# this script aims to set up the basic test env for the demo.
#

set -x

KAFKA_VER="2.13-2.8.0"
KAFKA_FOLDER="$(pwd)/kafka_${KAFKA_VER}"
KAFKA_TEST_TOPIC="test"
KAFKA_ADDR="localhost:9092"

# kill process id -1 is dangerous
KAFKA_PID=65535
ZK_PID=65535

# caution, do not modify the lock name
LOCK=/var/tmp/mykafka

trap clean EXIT

run_kafka() {
	if [ -f ${LOCK} ]; then
		echo "[ERROR] kafka is already running"
		exit 1
	fi
	touch ${LOCK}
	if [ -d ${KAFKA_FOLDER} ]; then
		cd ${KAFKA_FOLDER}
		TEST_RUN_TAG=$(date +%F-%H-%M)
		bin/zookeeper-server-start.sh config/zookeeper.properties 2>&1 > ${TEST_RUN_TAG}_zookeeper.txt &
		ZK_PID=$!
		sleep 10
		bin/kafka-server-start.sh config/server.properties 2>&1 > ${TEST_RUN_TAG}_kafka.txt &
		KAFKA_PID=$!
		echo "[INFO] Kafka service up: [zk_pid: ${ZK_PID}, kafka_pid: ${KAFKA_PID}"
	else
		echo "[ERROR] Kafka folder does not exist!"
		exit 1
	fi
}

create_kafka_topic() {
	cd ${KAFKA_FOLDER}
	bin/kafka-topics.sh --create --topic ${KAFKA_TEST_TOPIC} --bootstrap-server ${KAFKA_ADDR} 2>&1 > /dev/null
}

wait_zk_kafka() {
	wait ${ZK_PID}
	wait ${KAFKA_PID}
}

clean() {
	if [ ${KAFKA_PID} != 65535 ]; then
		kill -9 ${KAFKA_PID}
		rm ${LOCK}
	fi
	if [ ${ZK_PID} != 65535 ]; then
		kill -9 ${ZK_PID}
	fi
}

main() {

	run_kafka

	create_kafka_topic

	wait_zk_kafka

}

main $@
