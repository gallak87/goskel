// connects to grpcwebproxy which is listening on port 8080

import { GetUserRequest, GetUserResponse } from '@app/api/user_pb';
import { UserClient } from '@app/api/UserServiceClientPb';

const clientUser = new UserClient('http://localhost:8080');

export default class UserService {
  private readonly client: UserClient;

  constructor() {
    this.client = clientUser;
  }

  public async getUser(id: string): Promise<GetUserResponse> {
    const req = new GetUserRequest();
    req.setId(id);
    try {
      // TODO: metadata for auth tokens and such
      const metadata = null;
      return await this.client.getUser(req, metadata);
    } catch (e) {
      console.log(`getUser: GRPC error: ${e}`);
      throw e;
    }
  }
}
