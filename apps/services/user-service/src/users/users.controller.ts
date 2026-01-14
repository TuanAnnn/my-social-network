import { Controller, Post, Body, BadRequestException } from '@nestjs/common';
import { UsersService } from './users.service';
import { CreateUserDto } from './dto/create-user.dto'; // Import DTO

@Controller('api/users') // 
export class UsersController {
  constructor(private readonly usersService: UsersService) {}

  @Post('register')
  async register(@Body() body: CreateUserDto) { // <--- Dùng Class DTO thay vì object thường
    console.log('Body nhận được:', body);
    const existingUser = await this.usersService.findOne(body.email);
    if (existingUser) {
      throw new BadRequestException('Email này đã được sử dụng!');
    }

    const newUser = await this.usersService.createUser(body);

    const { password, ...result } = newUser;
    return {
      message: 'Đăng ký thành công',
      user: result
    };
  }
}