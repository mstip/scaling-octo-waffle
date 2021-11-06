import { IsNotEmpty, IsNumber, IsOptional, IsString } from 'class-validator';

export class CreateAssetDto {

    @IsString()
    @IsNotEmpty()
    name: string;

    @IsString()
    @IsOptional()
    purpose: string;

    @IsString()
    @IsOptional()
    owner: string;

    @IsString()
    @IsOptional()
    location: string;

    @IsNumber()
    @IsOptional()
    parentElementId: number;
}
