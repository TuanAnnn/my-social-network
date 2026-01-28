import {
  Controller,
  Post,
  Body,
  BadRequestException,
  Query,
  Get,
  Param,
  Patch,
} from "@nestjs/common";
import { UsersService } from "./users.service";
import { CreateUserDto } from "./dto/create-user.dto"; // Import DTO
import { UpdateUserDto } from "./dto/update-user.dto";

@Controller("api/users") //
export class UsersController {
  constructor(private readonly usersService: UsersService) {}

  @Post("register")
  async register(@Body() body: CreateUserDto) {
    console.log("Body nhận được:", body);
    const existingUser = await this.usersService.findOne(body.email);
    if (existingUser) {
      throw new BadRequestException("Email này đã được sử dụng!");
    }

    const newUser = await this.usersService.createUser(body);

    const { password, ...result } = newUser;
    return {
      message: "Đăng ký thành công",
      user: result,
    };
  }

  @Get()
  async getAllUser() {
    const allUser = await this.usersService.findAll();
    return allUser;
  }

  @Get(":id")
  async getUserById(@Param("id") id: number) {
    const user = await this.usersService.findOneById(id);
    return user;
  }

  @Patch(":id")
  async UpdateUserById(@Body() body: UpdateUserDto) {
    const user = await this.usersService.updateUser(body);
    return user;
  }
}
