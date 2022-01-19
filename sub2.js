const zmq = require("zeromq");

const sock = zmq.socket("sub");

sock.connect("tcp://127.0.0.1:8100");
sock.subscribe("185.107.232.99");
console.log("Subscriber connected to port 8100");

let counter = 0;
sock.on("message", function (topic, message) {
  console.log(
    "%d received a message related to: %s containing message:%s",
    counter++,
    topic.toString(),
    message.toString()
  );
});
