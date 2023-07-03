from moviepy.editor import VideoFileClip


def divide_video_into_clips(input_path, clip_duration=30):
    video = VideoFileClip(input_path)

    cnt = int(video.duration / clip_duration)

    for i in range(cnt):
        start_time = i * clip_duration
        end_time = (i + 1) * clip_duration

        clip = video.subclip(start_time, end_time)

        clip_output_path = f'clip_{i}.mp4'

        clip.write_videofile(clip_output_path, codec='libx264', audio_codec='aac')


video_path = '/dataset/videos/video.mp4'
divide_video_into_clips(video_path)
