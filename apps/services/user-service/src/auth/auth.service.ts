import { Injectable, UnauthorizedException } from '@nestjs/common';
import { UsersService } from '../users/users.service';
import { JwtService } from '@nestjs/jwt';
import * as bcrypt from 'bcrypt';

@Injectable()
export class AuthService {
  constructor(
    private usersService: UsersService,
    private jwtService: JwtService
  ) {}

  async signIn(email: string, pass: string) {
    // 1. Tìm user trong DB
    const user = await this.usersService.findOne(email);
    
    // 2. Nếu không có user -> Lỗi
    if (!user) {
      throw new UnauthorizedException('Email hoặc mật khẩu không đúng');
    }

    // 3. So sánh mật khẩu (pass nhập vào vs pass đã hash trong DB)
    const isMatch = await bcrypt.compare(pass, user.password);
    if (!isMatch) {
      throw new UnauthorizedException('Email hoặc mật khẩu không đúng');
    }

    // 4. Tạo Token (Payload chứa thông tin không nhạy cảm)
    const payload = { sub: user.id, email: user.email, name: user.name };
    
    return {
      access_token: await this.jwtService.signAsync(payload),
      user: {
        id: user.id,
        email: user.email,
        name: user.name
      }
    };
  }
}