import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { TypeOrmModule } from '@nestjs/typeorm';
import { UserModule } from './user/user.module';
import { AuthModule } from './auth/auth.module';
import { ConfigModule } from '@nestjs/config';
import { AssetModule } from './asset/asset.module';

@Module({
    imports: [
        TypeOrmModule.forRoot(),
        UserModule,
        AuthModule,
        ConfigModule.forRoot(),
        AssetModule
    ],
    controllers: [AppController],
    providers: [AppService],
})
export class AppModule {
}
