import { Client } from 'discord.js';

export enum CommandRegistrationResult {
    Success = 'Success',
    Error = 'Error'
}

export interface ICommand {
    name: string;
    register: (client: Client) => CommandRegistrationResult;
}
