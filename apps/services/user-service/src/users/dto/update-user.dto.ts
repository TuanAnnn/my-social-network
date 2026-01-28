import { IsEmail, IsNotEmpty, IsString, MinLength } from 'class-validator';

export class UpdateUserDto {
  @IsEmail({}, { message: 'Email không đúng định dạng' })
  email: string;

  @IsString({ message: 'Tên phải là chuỗi' })
  name: string;

  @IsString({ message: 'Avatar phải là chuỗi' })
  avatar: string;
}