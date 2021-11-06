import { IsEmail, IsString, MaxLength, MinLength } from "class-validator";

export class CreateUserDto {

    @MinLength(3)
    @MaxLength(256)
    @IsString()
    userName: string;

    @MinLength(3)
    @MaxLength(256)
    @IsEmail()
    email: string;

    @MinLength(3)
    @MaxLength(256)
    @IsString()
    password: string;
}
