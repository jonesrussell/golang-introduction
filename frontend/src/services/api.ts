import axios from 'axios';
import type { Tutorial, TutorialMetadata, Section, Exercise } from '../types/tutorial';
import type { Progress, ExecutionResult } from '../types/progress';

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
    // Use new path-based endpoint
    const response = await api.get<Tutorial>(`/tutorials/${id}`);
    return response.data;
  },

  async getTutorialSections(id: string): Promise<Section[]> {
    // Use new path-based endpoint
    const response = await api.get<Section[]>(`/tutorials/${id}/sections`);
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
    // Use new path-based endpoint
    const response = await api.get<Exercise[]>(`/exercises/${tutorialId}`);
    return response.data;
  },
};

