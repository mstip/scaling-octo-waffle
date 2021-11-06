import {Injectable} from '@nestjs/common';
import {CreateMessageDto} from './dto/create-message.dto';
import {UpdateMessageDto} from './dto/update-message.dto';
import {InjectRepository} from "@nestjs/typeorm";
import {Message} from "./entities/message.entity";
import {DeleteResult, Repository} from "typeorm";

@Injectable()
export class MessageService {

    constructor(
        @InjectRepository(Message)
        private messageRepository: Repository<Message>
    ) {
    }

    create(createMessageDto: CreateMessageDto): Promise<Message> {
        return this.messageRepository.save(createMessageDto);
    }

    findAll(): Promise<Message[]> {
        return this.messageRepository.find();
    }

    findOne(id: number): Promise<Message> {
        return this.messageRepository.findOne(id);
    }

    update(id: number, updateMessageDto: UpdateMessageDto) {
        updateMessageDto.id = id;
        return this.messageRepository.save(updateMessageDto);
    }

    remove(id: number): Promise<DeleteResult> {
        return this.messageRepository.delete(id);
    }
}
