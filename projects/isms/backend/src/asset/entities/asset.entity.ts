import { Column, Entity, PrimaryGeneratedColumn } from 'typeorm';

@Entity()
export class Asset {
    @PrimaryGeneratedColumn()
    id: number;

    @Column()
    name: string;

    @Column({nullable: true})
    purpose: string;

    @Column({nullable: true})
    owner: string;

    @Column({nullable: true})
    location: string;

    @Column({nullable: true})
    parentElementId: number;
}
