import {Module} from '@nestjs/common';
import {AppController} from './app.controller';
import {AppService} from './app.service';
import {MessageModule} from './message/message.module';
import {TypeOrmModule} from '@nestjs/typeorm';
import {Message} from "./message/entities/message.entity";

@Module({
    imports: [
        MessageModule,
        TypeOrmModule.forRoot({type: 'sqlite', database: ':memory:' , synchronize: true, entities:[Message]})
    ],
    controllers: [AppController],
    providers: [AppService],
})
export class AppModule {
}
