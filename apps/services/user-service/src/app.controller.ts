import { Controller, Get } from '@nestjs/common';

@Controller('/')
export class AppController {
  
  @Get()
  getUsers() {
    return {
      service: "NestJS User Service",
      timestamp: new Date(),
      data: [
        { id: 1, name: "Nest Admin", role: "SuperAdmin" },
        { id: 2, name: "Nest User", role: "Member" }
      ]
    };
  }
}