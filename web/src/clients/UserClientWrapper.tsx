import { UserServiceClient, ServiceError } from '../api/user_pb_service';
import { GetUserRequest, GetUserResponse } from '../api/user_pb'

// connects to grpcwebproxy which is listening on port 8080
// TODO: configurize eventually..
const client = new UserServiceClient("http://localhost:8080");

// Wrapper class around grpc request/response objects
const UserClientWrapper = () => {
  // TODO:
  // const getUser = (id: string): GetUserResponse =>  {
  //   const req = new GetUserRequest();
  //   req.setId("1");
  //   let name:string|undefined = 'unknown name';
  //   client.getUser(req,
  //     (error: ServiceError | null, response: GetUserResponse | null) => {
  //       console.log(error);
  //       console.log(response);
  //       name = response?.getName();
  //       console.log(name);
  // }
}