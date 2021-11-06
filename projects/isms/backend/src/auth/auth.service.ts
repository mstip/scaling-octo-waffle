import { Injectable } from '@nestjs/common';
import { UserService } from '../user/user.service';
import * as bcrypt from 'bcrypt';
import { User } from '../user/entities/user.entity';
import { JwtService } from '@nestjs/jwt';

@Injectable()
export class AuthService {
    constructor(private userService: UserService, private jwtService: JwtService) {
    }

    async validateUser(username: string, password: string): Promise<User> {
        const user = await this.userService.findByUserName(username);
        if (!user) {
            return null;
        }

        if (!await bcrypt.compare(password, user.passwordHash)) {
            return null;
        }

        return user;
    }

    async login(user: any) {
        const payload = {username: user.username, sub: user.userId};
        return {
            access_token: this.jwtService.sign(payload),
        };
    }

}