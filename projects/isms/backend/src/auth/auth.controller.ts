import { ClassSerializerInterceptor, Controller, Post, Request, UseGuards, UseInterceptors } from '@nestjs/common';
import { LocalAuthGuard } from './local-auth.guard';
import { AuthService } from './auth.service';

@Controller('auth')
export class AuthController {
    constructor(private authService: AuthService) {
    }

    @UseInterceptors(ClassSerializerInterceptor)
    @UseGuards(LocalAuthGuard)
    @Post('login')
    async login(@Request() req) {
        return this.authService.login(req.user);
    }

}
