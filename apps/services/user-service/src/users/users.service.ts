import { Injectable } from "@nestjs/common";
import { PrismaService } from "../prisma/prisma.service";
import { Prisma } from "@prisma/client";
import * as bcrypt from "bcrypt"; // Import thư viện
import { updateUserData } from "./type";

@Injectable()
export class UsersService {
  constructor(private prisma: PrismaService) {}

  async createUser(data: Prisma.UserCreateInput) {
    const salt = await bcrypt.genSalt(10);
    const hashedPassword = await bcrypt.hash(data.password, salt);

    return this.prisma.user.create({
      data: {
        ...data,
        password: hashedPassword,
      },
    });
  }

  async findOne(email: string) {
    return this.prisma.user.findUnique({ where: { email } });
  }

  async findOneById(id: number) {
    return this.prisma.user.findUnique({ where: { id } });
  }

  async findAll() {
    return this.prisma.user.findMany({});
  }

  async updateUser(data: updateUserData) {
    const updatedUser = await this.prisma.user.update({
      where: { email: data.email },
      data: {
        email: data.email,
        avatar: data.avatar,
        name: data.name,
      },
    });
    return updatedUser;
  }
}
