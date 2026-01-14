import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  // Chạy ở port 3000
  await app.listen(3000);
  console.log(`NestJS User Service is running on port 3000`);
}
bootstrap();