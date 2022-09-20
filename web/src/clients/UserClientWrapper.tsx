// connects to grpcwebproxy which is listening on port 8080

import { UserServicePromiseClient } from "../api/user_grpc_web_pb";
import { GetUserRequest, GetUserResponse } from "../api/user_pb";

// TODO: configurize eventually..
const client = new UserServicePromiseClient("http://localhost:8080");

// Wrapper class around grpc request/response objects
export default class UserClient {
  public async getUser(id: string): Promise<GetUserResponse> {
    const req = new GetUserRequest();
    req.setId(id);
    try {
      return await client.getUser(req);
    } catch (e) {
      console.log(`getUser: GRPC error: ${e}`)
      throw e;
    }
  }
}
