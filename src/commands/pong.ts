import { ICommand, CommandRegistrationResult } from './../interfaces';
import { Client, Message } from 'discord.js';

const register = (client: Client) => {
    client.on('message', (message: Message) => {
        if (message.content.toLowerCase() === 'pong') {
            message.reply('ping');
        }
    });
    return CommandRegistrationResult.Success;
};

const command: ICommand = {
    name: 'Pong',
    register
};

export default command;
