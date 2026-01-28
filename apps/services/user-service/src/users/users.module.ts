import { Module } from '@nestjs/common';
import { UsersService } from './users.service';
import { UsersController } from './users.controller';
import { PrismaModule } from '../prisma/prisma.module';

@Module({
  controllers: [UsersController],
  providers: [UsersService],
  imports: [PrismaModule], 
  exports: [UsersService], // <--- BẮT BUỘC PHẢI CÓ DÒNG NÀY
})
export class UsersModule {}