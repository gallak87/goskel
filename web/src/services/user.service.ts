import { UserClient } from "../api/UserServiceClientPb";
import { GetUserResponse, GetUserRequest } from "../api/user_pb";

// connects to grpcwebproxy which is listening on port 8080
// TODO: configurize eventually..
const client = new UserClient("http://localhost:8080");

export default class UserService {
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
