
//  npx webpack  .\index.js 
// grpcwebproxy  --backend_addr=localhost:9000  --run_tls_server=false --allow_all_origins

const {DetailReq} = require('./blog/blog_pb.js');
const {BlogClient} = require('./blog/blog_grpc_web_pb.js');
var client = new BlogClient('http://localhost:8080');
var request = new DetailReq();
request.setId(1);
client.detail(request, {}, (err, response) => {
  console.log(response.toString());
});