import axios from 'axios';
import type { Tutorial, TutorialMetadata, Section, Exercise, Progress, ExecutionResult } from '../types';

const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080/api';

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

export const tutorialApi = {
  async listTutorials(): Promise<TutorialMetadata[]> {
    const response = await api.get<TutorialMetadata[]>('/tutorials');
    return response.data;
  },

  async getTutorial(id: string): Promise<Tutorial> {
    const response = await api.get<Tutorial>('/tutorial', { params: { id } });
    return response.data;
  },

  async getTutorialSections(id: string): Promise<Section[]> {
    const response = await api.get<Section[]>('/tutorial/sections', { params: { id } });
    return response.data;
  },
};

export const executionApi = {
  async executeCode(code: string): Promise<ExecutionResult> {
    const response = await api.post<ExecutionResult>('/execute', { code });
    return response.data;
  },
};

export const progressApi = {
  async getProgress(userId: string = 'default'): Promise<Progress> {
    const response = await api.get<Progress>('/progress', { params: { userId } });
    return response.data;
  },

  async updateProgress(progress: Progress, userId: string = 'default'): Promise<void> {
    await api.post('/progress', progress, { params: { userId } });
  },

  async markSectionComplete(tutorialId: string, sectionId: string, userId: string = 'default'): Promise<void> {
    await api.post('/progress/section', { tutorialId, sectionId }, { params: { userId } });
  },
};

export const exerciseApi = {
  async getExercises(tutorialId: string): Promise<Exercise[]> {
    const response = await api.get<Exercise[]>('/exercises', { params: { tutorialId } });
    return response.data;
  },
};

