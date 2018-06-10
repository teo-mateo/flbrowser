package rtorrent

import "github.com/teo-mateo/flbrowser/dto"

func GetTorrents() ([]*dto.RTRTorrentInfo, error){
	ids, err := RPC_download_list()
	if err != nil{
		return nil, err
	}
	result := make([]*dto.RTRTorrentInfo, 0)

	for _, id := range ids{
		crt := dto.RTRTorrentInfo{
			ID: id,
		}

		result = append(result, &crt)

		err = getRTRInfo(&crt)
		if err != nil{
			return nil, err
		}
	}

	return result, nil
}

func getRTRInfo(dto *dto.RTRTorrentInfo) error {

	var err error

	//get name
		dto.Name, err = RPC_id__string("d.name", dto.ID)
	if err != nil{
		return err
	}

	//get creation date
	dto.CreationDate, err = RPC_id__int("d.creation_date", dto.ID)
	if err != nil{
		return err
	}

	//get is open
	dto.IsOpen, err = RPC_id__bool("d.is_open", dto.ID)
	if err != nil{
		return err
	}

	//get is active
	dto.IsActive, err = RPC_id__bool("d.is_active", dto.ID)
	if err != nil{
		return err
	}

	//get is hash checked
	dto.IsHashChecked, err = RPC_id__bool("d.is_hash_checked", dto.ID)
	if err != nil{
		return err
	}

	//get is hash checking
	dto.IsHashChecking, err = RPC_id__bool("d.is_hash_checking", dto.ID)
	if err != nil{
		return err
	}

	//get is multi file
	dto.IsMultiFile, err = RPC_id__bool("d.is_multi_file", dto.ID)
	if err != nil{
		return err
	}

	//total downloaded
	dto.DownTotal, err = RPC_id__int("d.down.total", dto.ID)
	if err != nil{
		return err
	}

	//total uploaded
	dto.UpTotal, err = RPC_id__int("d.up.total", dto.ID)
	if err != nil{
		return err
	}

	//get directory
	dto.Directory, err = RPC_id__string("d.directory", dto.ID)
	if err != nil{
		return err
	}

	//get completed bytes
	dto.CompletedBytes, err = RPC_id__int("d.completed_bytes", dto.ID)
	if err != nil{
		return err
	}

	//get left bytes
	dto.LeftBytes, err = RPC_id__int("d.left_bytes", dto.ID)
	if err != nil{
		return err
	}

	//get size files
	dto.SizeFiles, err = RPC_id__int("d.size_files", dto.ID)
	if err != nil{
		return err
	}

	//get size bytes
	dto.SizeBytes, err = RPC_id__int("d.size_bytes", dto.ID)
	if err != nil{
		return err
	}



	return nil
}
