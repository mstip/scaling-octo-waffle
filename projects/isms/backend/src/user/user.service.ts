import { Injectable, OnApplicationBootstrap } from '@nestjs/common';
import { CreateUserDto } from './dto/create-user.dto';
import { InjectRepository } from '@nestjs/typeorm';
import { User } from './entities/user.entity';
import { Repository } from 'typeorm';
import * as bcrypt from 'bcrypt';

@Injectable()
export class UserService implements OnApplicationBootstrap {

    constructor(
        @InjectRepository(User)
        private usersRepository: Repository<User>,
    ) {
    }

    async onApplicationBootstrap(): Promise<void> {
        // TODO: check if database seeding is a better idea
        await this.createFirstUser();
    }

    async create(createUserDto: CreateUserDto): Promise<User> {
        //TODO: maybe move this to entity
        const newUser: User = new User({});
        newUser.userName = createUserDto.userName;
        newUser.email = createUserDto.email;
        newUser.passwordHash = await bcrypt.hash(createUserDto.password, 10);
        return this.usersRepository.save(newUser);
    }

    findAll(): Promise<User[]> {
        return this.usersRepository.find();
    }

    findByUserName(userName: string): Promise<User | undefined> {
        return this.usersRepository.findOne({userName});
    }

    async createFirstUser(): Promise<void> {
        const userCount = await this.usersRepository.count();
        if (userCount > 0) {
            return;
        }
        await this.create({
            userName: 'admin',
            password: 'admin',
            email: 'admin@admin.de'
        });
    }
}
