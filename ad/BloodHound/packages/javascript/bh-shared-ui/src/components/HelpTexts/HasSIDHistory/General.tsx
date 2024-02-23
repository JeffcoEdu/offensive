// Copyright 2023 Specter Ops, Inc.
//
// Licensed under the Apache License, Version 2.0
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

import { FC } from 'react';
import { typeFormat } from '../utils';
import { EdgeInfoProps } from '../index';
import { Typography } from '@mui/material';

const General: FC<EdgeInfoProps> = ({ sourceName, sourceType, targetName, targetType }) => {
    return (
        <>
            <Typography variant='body2'>
                The {typeFormat(sourceType)} {sourceName} has, in its SIDHistory attribute, the SID for the{' '}
                {typeFormat(targetType)} {targetName}.
            </Typography>
            <Typography variant='body2'>
                When a kerberos ticket is created for {sourceName}, it will include the SID for {targetName}, and
                therefore grant
                {sourceName} the same privileges and permissions as
                {targetName}.
            </Typography>
        </>
    );
};

export default General;