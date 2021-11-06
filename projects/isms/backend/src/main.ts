import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { ValidationPipe } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';

async function bootstrap() {
    const app = await NestFactory.create(AppModule);
    const configService = app.get(ConfigService);
    app.enableCors();
    if (configService.get('APP_ENV') !== 'prod') {
        app.setGlobalPrefix('api');
    }
    app.useGlobalPipes(new ValidationPipe());
    await app.listen(3000);
}

bootstrap();
