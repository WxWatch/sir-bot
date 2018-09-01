import { Client, Message } from 'discord.js';
import * as glob from 'glob';

import { logger } from './logger';
import { ICommand, CommandRegistrationResult, ICron } from './interfaces';

const client = new Client();
const token = process.env.DISCORD_TOKEN;

if (!token) {
    logger.error('DISCORD_TOKEN not set!');
    process.exit(1);
}

let commands: ICommand[] = [];
let crons: ICron[] = [];

export const registeredCommands = () => {
    return commands;
};

const registerCrons = (client: Client) => {
    client.on('ready', async () => {
        crons = await loadCrons();
        logger.info(`Registering ${crons.length} crons...`);
        for (const cron of crons) {
            setInterval(() => {
                cron.action(client);
            }, cron.interval);
        }
    });
};

const loadCrons = async () => {
    const crons = [];
    const files = glob.sync('./cron/**/*', {
        ignore: './cron/**/*.map',
        cwd: './dist'
    });
    logger.info(`Loading ${files.length} crons`);
    for (const file of files) {
        const command = await import(file);
        crons.push(command.default);
    }

    return crons;
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
    const files = glob.sync('./commands/**/*', {
        ignore: './commands/**/*.map',
        cwd: './dist'
    });
    logger.info(`Loading ${files.length} commands`);
    for (const file of files) {
        const command = await import(file);
        commands.push(command.default);
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
registerCrons(client);
client.login(token);
