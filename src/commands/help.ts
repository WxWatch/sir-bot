import { ICommand, CommandRegistrationResult } from './../interfaces';
import { Client } from 'discord.js';

const command: ICommand = {
    name: 'Help',
    register: (client: Client): CommandRegistrationResult => {
        return CommandRegistrationResult.Success;
    }
};

export default command;
