import { Column, Entity, PrimaryGeneratedColumn } from "typeorm";
import { Exclude } from "class-transformer";

@Entity()
export class User {

    @PrimaryGeneratedColumn()
    id: number;

    @Column({
        unique: true,
    })
    email: string;

    @Column({
        unique: true,
    })
    userName: string;

    @Column()
    @Exclude()
    passwordHash: string;


    constructor(partial: Partial<User>) {
        Object.assign(this, partial);
    }

}
