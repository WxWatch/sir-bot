import { ICommand, CommandRegistrationResult } from './../interfaces';
import { Client } from 'discord.js';

const register = (client: Client): CommandRegistrationResult => {
    return CommandRegistrationResult.Success;
};

const command: ICommand = {
    name: 'Help',
    register: (client: Client) => {
        return CommandRegistrationResult.Success;
    }
};

export default command;
