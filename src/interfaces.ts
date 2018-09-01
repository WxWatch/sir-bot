import { Client } from 'discord.js';

export enum CommandRegistrationResult {
    Success = 'Success',
    Error = 'Error'
}

export interface ICommand {
    name: string;
    register: (client: Client) => CommandRegistrationResult;
}

export interface ICron {
    name: string;
    interval: number;
    action: (client: Client) => void;
}
