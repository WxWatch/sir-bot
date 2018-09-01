import { Client, Message } from 'discord.js';
import * as glob from 'glob';

import { logger } from './logger';
import { ICommand, CommandRegistrationResult } from './interfaces';

const client = new Client();
const token = process.env.DISCORD_TOKEN || 'NDg1MTAxMTM1MDI4Mjg5NTc3.DmrpRA.PsFvFqLpWvjkeM0GU5g7vvMvOlg';

let commands: ICommand[] = [];

export const registeredCommands = () => {
    return commands;
};

const registerListeners = (client: Client) => {
    client.on('ready', async () => {
        logger.info(`Logged in as ${client.user.tag}!`);
        commands = await loadCommands();
        registerCommands(commands);
    });
};

const loadCommands = async (): Promise<ICommand[]> => {
    const commands = [];
    const files = glob.sync('./commands/*.js', {
        cwd: './dist'
    });
    logger.info(`Loading ${files.length} commands`);
    for (const file of files) {
        const command: ICommand = await import(file);
        commands.push(command);
    }

    return commands;
};

const registerCommands = async (commands: ICommand[]) => {
    logger.info(`Registering ${commands.length} commands...`);
    for (const command of commands) {
        const result = command.register(client);
        logger.info(`${command.name}: ${result}`);
    }
};

registerListeners(client);
client.login(token);
