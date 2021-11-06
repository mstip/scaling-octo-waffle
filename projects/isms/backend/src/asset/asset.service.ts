import { Injectable } from '@nestjs/common';
import { CreateAssetDto } from './dto/create-asset.dto';
import { UpdateAssetDto } from './dto/update-asset.dto';
import { InjectRepository } from '@nestjs/typeorm';
import { DeleteResult, Repository } from 'typeorm';
import { Asset } from './entities/asset.entity';

@Injectable()
export class AssetService {
    constructor(
        @InjectRepository(Asset)
        private assetRepository: Repository<Asset>,
    ) {
    }

    create(createAssetDto: CreateAssetDto): Promise<Asset> {
        return this.assetRepository.save(createAssetDto);
    }

    findAll(): Promise<Asset[]> {
        return this.assetRepository.find();
    }

    findOne(id: number): Promise<Asset | undefined> {
        return this.assetRepository.findOne(id);
    }

    async update(id: number, updateAssetDto: UpdateAssetDto): Promise<Asset> {
        try {
            await this.assetRepository.findOneOrFail(id);
        } catch (err) {
            return null;
        }
        return this.assetRepository.save({id, ...updateAssetDto});
    }

    remove(id: number): Promise<DeleteResult> {
        return this.assetRepository.delete(id);
    }
}
