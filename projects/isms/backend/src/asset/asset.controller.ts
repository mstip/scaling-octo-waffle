import { Body, Controller, Delete, Get, NotFoundException, Param, Patch, Post, UseGuards } from '@nestjs/common';
import { AssetService } from './asset.service';
import { CreateAssetDto } from './dto/create-asset.dto';
import { UpdateAssetDto } from './dto/update-asset.dto';
import { JwtAuthGuard } from '../auth/jwt-auth.guard';

@Controller('asset')
export class AssetController {
    constructor(private readonly assetService: AssetService) {
    }

    @UseGuards(JwtAuthGuard)
    @Post()
    create(@Body() createAssetDto: CreateAssetDto) {
        return this.assetService.create(createAssetDto);
    }

    @UseGuards(JwtAuthGuard)
    @Get()
    findAll() {
        return this.assetService.findAll();
    }

    @UseGuards(JwtAuthGuard)
    @Get(':id')
    async findOne(@Param('id') id: string) {
        const asset = await this.assetService.findOne(+id);
        if (!asset) {
            throw new NotFoundException();
        }
        return asset;
    }

    @UseGuards(JwtAuthGuard)
    @Patch(':id')
    async update(@Param('id') id: string, @Body() updateAssetDto: UpdateAssetDto) {
        const asset = await this.assetService.update(+id, updateAssetDto);
        if (!asset) {
            throw new NotFoundException();
        }
        return asset;
    }

    @UseGuards(JwtAuthGuard)
    @Delete(':id')
    remove(@Param('id') id: string) {
        return this.assetService.remove(+id);
    }
}
