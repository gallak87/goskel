// connects to grpcwebproxy which is listening on port 8080

import { UserServiceClient } from "../api/UserServiceClientPb";
import { GetUserResponse, GetUserRequest } from "../api/user_pb";

// TODO: configurize eventually..
const client = new UserServiceClient("http://localhost:8080");

export default class UserClient {
  public async getUser(id: string): Promise<GetUserResponse> {
    const req = new GetUserRequest();
    req.setId(id);
    try {
      // TODO: metadata for auth tokens and such
      const metadata = null;
      return await client.getUser(req, metadata);
    } catch (e) {
      console.log(`getUser: GRPC error: ${e}`)
      throw e;
    }
  }
}
