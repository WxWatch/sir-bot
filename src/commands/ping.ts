import { ICommand, CommandRegistrationResult } from './../interfaces';
import { Client, Message } from 'discord.js';

const command: ICommand = {
    name: 'Ping',
    register: (client: Client) => {
        client.on('message', (message: Message) => {
            if (message.content === 'ping') {
                message.reply('pong');
            }
        });
        return CommandRegistrationResult.Success;
    }
};

export default command;
